import { createBrowserRouter } from "react-router-dom";

import Home from "./Home";
import Login from "./Login";

export const router = createBrowserRouter([
  {
    index: true,
    element: <Home />,
  },
  {
    path: "/login",
    element: <Login />,
  },
]);
export default router;
