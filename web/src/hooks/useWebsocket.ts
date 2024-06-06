import { MutableRefObject, useEffect, useRef, useState } from "react";
import { baseURL, token } from "@/network/api.ts";

import useInterval from "@hooks/useInterval.ts";

const useWebsocket = (
  url: string | null,
  handler: (msg: MessageEvent) => void,
): [MutableRefObject<WebSocket | null>, boolean, Event | null] => {
  const conn = useRef<WebSocket | null>(null);
  const [connected, setConnected] = useState(false);
  const [error, setError] = useState<Event | null>(null);
  const callback = useRef(handler);

  const handleConnect = () => {
    conn.current?.close();
    if (url === null) return;
    conn.current = new WebSocket(
      `${location.protocol === "https:" ? "wss" : "ws"}://${location.host + baseURL + url}`,
      token ?? undefined,
    );
    conn.current.onopen = () => {
      setConnected(true);
      setError(null);
    };
    conn.current.onerror = (err) => setError(err);
    conn.current.onmessage = (ev) => callback.current(ev);
    conn.current.onclose = () => setConnected(false);
  };

  useInterval(handleConnect, url !== null && !connected ? 4000 : null);

  useEffect(() => {
    conn.current?.close();
    setConnected(false);
    setError(null);
    if (url !== null) handleConnect();

    return () => {
      conn.current?.close();
    };
  }, [url]);
  useEffect(() => {
    callback.current = handler;
  }, [handler]);

  return [conn, connected, error];
};
export default useWebsocket;
