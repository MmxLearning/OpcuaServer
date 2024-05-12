import { FC, useState } from "react";
import toast from "react-hot-toast";

import {
  PickerValidDate,
  DateTimePicker,
  LocalizationProvider,
} from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";

import { Container, Stack, Divider, TextField, Button } from "@mui/material";
import { HorizontalRule } from "@mui/icons-material";

import { useApi } from "@/network/api.ts";

export const Home: FC = () => {
  const [name, setName] = useState("");
  const [nodeID, setNodeID] = useState("");
  const [startAt, setStartAt] = useState<PickerValidDate | null>(null);
  const [endAt, setEndAt] = useState<PickerValidDate | null>(null);

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
            padding: "3rem 2rem",
          }}
          spacing={2.5}
        >
          <Stack
            direction={{ md: "row", xs: "column" }}
            sx={{
              "&>.MuiTextField-root": {
                flex: 1,
              },
            }}
          >
            <TextField label={"Name"} />
            <div style={{ width: "2rem", height: "1.5rem" }} />
            <TextField label={"NodeID"} />
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

            <Button
              variant={"contained"}
              sx={{
                my: 0.1,
                flex: 1,
              }}
            >
              Search
            </Button>
          </Stack>
        </Stack>
      </Container>
    </LocalizationProvider>
  );
};
export default Home;
