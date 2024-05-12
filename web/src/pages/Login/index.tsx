import { FC, useState } from "react";
import { useNavigate } from "react-router-dom";
import toast from "react-hot-toast";

import {
  TextField,
  Stack,
  Paper,
  Container,
  Button,
  Typography,
} from "@mui/material";

import { api, setToken } from "@/network/api.ts";

export const Login: FC = () => {
  const nav = useNavigate();

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const onLogin = async () => {
    if (!username || !password) {
      toast.error("Username and password are required");
      return;
    }
    if (username.length > 15) {
      toast.error("Username must be less than 15 characters");
      return;
    }
    if (password.length > 72 || password.length < 16) {
      toast.error("Password must be between 16 and 72 characters");
      return;
    }

    try {
      const {
        data: { token },
      } = await api.post<{ token: string }>("public/login", {
        username,
        password,
      });
      setToken(token);
      toast.success("登录成功");
      nav("/");
    } catch (err: any) {
      if (err.msg) toast.error(err.msg);
    }
  };

  return (
    <Container
      sx={{
        height: "100%",
      }}
    >
      <Stack
        alignItems={"center"}
        justifyContent={"center"}
        sx={{
          height: "100%",
        }}
      >
        <Stack
          component={Paper}
          elevation={4}
          spacing={2}
          alignItems={"center"}
          sx={{
            padding: "3rem 4rem",
          }}
        >
          <Typography variant={"h4"}>OPC UA Sync Server</Typography>
          <Stack
            alignItems={"center"}
            spacing={2}
            pt={3}
            sx={{
              width: "16.5rem",
              maxWidth: "100%",
            }}
          >
            <TextField
              fullWidth
              label={"Username"}
              value={username}
              onChange={(ev) => setUsername(ev.target.value)}
            />
            <TextField
              fullWidth
              label={"Password"}
              type={"password"}
              value={password}
              onChange={(ev) => setPassword(ev.target.value)}
            />
            <div />
            <Button variant={"contained"} fullWidth onClick={onLogin}>
              Login
            </Button>
          </Stack>
        </Stack>
      </Stack>
    </Container>
  );
};
export default Login;
