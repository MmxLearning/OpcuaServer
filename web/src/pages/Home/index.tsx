import { FC, useState } from "react";
import useWebsocket from "@hooks/useWebsocket.ts";
import JsonView from "@uiw/react-json-view";

import {
  PickerValidDate,
  DateTimePicker,
  LocalizationProvider,
} from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DataGrid } from "@mui/x-data-grid";

import { LoadingButton } from "@mui/lab";
import {
  Container,
  Stack,
  TextField,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Grid,
  Card,
  Typography,
  CardContent,
} from "@mui/material";
import { HorizontalRule, InfoOutlined, Stream } from "@mui/icons-material";

import { useApi } from "@/network/api.ts";

export const Home: FC = () => {
  const [name, setName] = useState("");
  const [nodeID, setNodeID] = useState("");
  const [startAt, setStartAt] = useState<PickerValidDate | null>(null);
  const [endAt, setEndAt] = useState<PickerValidDate | null>(null);

  const [onSearching, setOnSearching] = useState("");

  const { isLoading, data } = useApi<Opcua.SearchResult[]>(
    onSearching ? `user/opcua/search?${onSearching}` : null,
  );

  const [isViewingDetail, setIsViewingDetail] = useState(false);
  const [onViewDetail, setOnViewDetail] = useState<Opcua.SearchResult | null>(
    null,
  );

  const [isViewingRdp, setIsViewingRdp] = useState(false);
  const { isLoading: isRdpDataLoading, data: rdpData } = useApi<Rdp.Info[]>(
    isViewingRdp ? `user/rdp/` : null,
  );
  const [onViewStream, setOnViewStream] = useState<Rdp.OnViewStream | null>(
    null,
  );

  const [streamFrame, setStreamFrame] = useState<string | undefined>(undefined);
  useWebsocket(
    onViewStream
      ? `user/rdp/stream?name=${encodeURI(onViewStream.name)}`
      : null,
    (ev) => setStreamFrame(window.URL.createObjectURL(ev.data)),
  );

  const handleSearch = () => {
    setOnSearching(
      new URLSearchParams({
        name,
        nodeID,
        startAt: startAt?.unix().toString() ?? "0",
        endAt: endAt?.unix().toString() ?? "0",
      }).toString(),
    );
  };

  return (
    <LocalizationProvider dateAdapter={AdapterDayjs}>
      <Container
        sx={{
          width: "100%",
          height: "100vh",
          overflow: "hidden",
        }}
      >
        <Stack
          alignItems={"stretch"}
          sx={{
            boxSizing: "border-box",
            padding: "3rem 2rem",
            height: "100%",
          }}
          spacing={2.5}
        >
          <Stack
            direction={"row"}
            sx={{
              width: "100%",
            }}
          >
            <LoadingButton
              startIcon={<Stream />}
              variant={"outlined"}
              onClick={() => setIsViewingRdp(true)}
              loading={isRdpDataLoading}
            >
              Remote Devices
            </LoadingButton>
          </Stack>
          <Stack
            direction={{ md: "row", xs: "column" }}
            sx={{
              "&>.MuiTextField-root": {
                flex: 1,
              },
            }}
          >
            <TextField
              label={"Name"}
              value={name}
              onChange={(ev) => setName(ev.target.value)}
            />
            <div style={{ width: "2rem", height: "1.5rem" }} />
            <TextField
              label={"NodeID"}
              value={nodeID}
              onChange={(ev) => setNodeID(ev.target.value)}
            />
          </Stack>
          <Stack
            direction={{ md: "row", xs: "column" }}
            justifyContent={"space-between"}
          >
            <Stack direction={"row"} alignItems={"center"}>
              <DateTimePicker
                disableFuture
                value={startAt}
                maxDateTime={endAt?.clone()}
                onChange={(ev) => setStartAt(ev)}
              />
              <HorizontalRule
                sx={{
                  color: "text.secondary",
                  fontSize: 20,
                  mx: 1.6,
                }}
              />
              <DateTimePicker
                value={endAt}
                minDateTime={startAt?.clone()}
                onChange={(ev) => setEndAt(ev)}
              />
            </Stack>

            <div style={{ width: "2.5rem", height: "1.5rem" }} />

            <LoadingButton
              variant={"contained"}
              loading={isLoading}
              sx={{
                my: 0.1,
                flex: 1,
              }}
              onClick={handleSearch}
            >
              Search
            </LoadingButton>
          </Stack>

          <DataGrid
            rows={data ?? []}
            disableColumnSelector
            disableDensitySelector
            columns={[
              {
                field: "id",
                headerName: "ID",
                minWidth: 80,
                flex: 1,
              },
              {
                field: "name",
                headerName: "Name",
                minWidth: 150,
                flex: 3,
              },
              {
                field: "created_at",
                headerName: "Created At",
                minWidth: 160,
                flex: 5,
                renderCell: (params) => {
                  return new Date(params.value * 1000).toLocaleString();
                },
              },
              {
                field: "timestamp",
                headerName: "Timestamp",
                minWidth: 160,
                flex: 5,
                renderCell: (params) => {
                  return new Date(params.value * 1000).toLocaleString();
                },
              },
              {
                field: "action",
                headerName: "Action",
                sortable: false,
                filterable: false,
                disableColumnMenu: true,
                width: 160,
                renderCell: (params) => {
                  return (
                    <>
                      <Button
                        variant="outlined"
                        startIcon={<InfoOutlined />}
                        onClick={() => {
                          setIsViewingDetail(true);
                          setOnViewDetail(params.row);
                        }}
                      >
                        Detail
                      </Button>
                    </>
                  );
                },
              },
            ]}
            initialState={{
              pagination: {
                paginationModel: { page: 0, pageSize: 30 },
              },
            }}
            pageSizeOptions={[30, 60, 100]}
          />
        </Stack>
      </Container>

      <Dialog
        open={isViewingDetail}
        onClose={() => setIsViewingDetail(false)}
        scroll={"paper"}
        fullWidth
        maxWidth={"lg"}
      >
        <DialogTitle>
          {onViewDetail?.name} {onViewDetail?.id}
        </DialogTitle>
        <DialogContent dividers>
          <JsonView
            value={onViewDetail ? JSON.parse(onViewDetail.data) : undefined}
          />
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setIsViewingDetail(false)}>Close</Button>
        </DialogActions>
      </Dialog>

      <Dialog
        open={!!rdpData && isViewingRdp}
        onClose={() => setIsViewingRdp(false)}
        scroll={"paper"}
        fullWidth
        maxWidth={"md"}
      >
        <DialogTitle>Remote Devices</DialogTitle>
        <DialogContent dividers>
          <Grid container spacing={2}>
            {rdpData?.map((item) => (
              <Grid key={item.name} item md={4} sm={6}>
                <Card
                  elevation={2}
                  onClick={() =>
                    setOnViewStream({
                      name: item.name,
                      frame_rate: item.frame_rate,
                    })
                  }
                >
                  <CardContent
                    sx={{
                      userSelect: "none",
                    }}
                  >
                    <Typography variant={"h5"}>{item.name}</Typography>
                    <Typography variant={"body2"} color={"text.secondary"}>
                      {item.desc}
                    </Typography>
                  </CardContent>
                </Card>
              </Grid>
            ))}
          </Grid>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setIsViewingRdp(false)}>Close</Button>
        </DialogActions>
      </Dialog>

      <Dialog
        open={!!onViewStream}
        onClose={() => setOnViewStream(null)}
        scroll={"paper"}
        fullWidth
        maxWidth={"xl"}
        sx={{
          "& .MuiDialog-paper": {
            height: "100%",
          },
        }}
      >
        <DialogTitle>Remote Desktop {onViewStream?.name}</DialogTitle>
        <DialogContent dividers>
          <Stack
            justifyContent={"center"}
            alignItems={"center"}
            sx={{
              height: "100%",
              width: "100%",
              "& img": {
                maxWidth: "100%",
                maxHeight: "100%",
              },
            }}
          >
            <img src={streamFrame} />
          </Stack>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOnViewStream(null)}>Close</Button>
        </DialogActions>
      </Dialog>
    </LocalizationProvider>
  );
};
export default Home;
