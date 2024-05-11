import { FC } from "react";
import useMount from "@hooks/useMount.ts";

import { RouterProvider } from "react-router-dom";
import router from "@/pages/router.tsx";

import { token, goLogin } from "@/network/api.ts";

export const App: FC = () => {
  useMount(() => {
    if (import.meta.env.MODE === "production") {
      if (!token) goLogin();
    }
  });

  return <RouterProvider router={router} />;
};
export default App;
