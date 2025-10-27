import { useSuspenseQuery } from "@tanstack/react-query";
import { authQueryOptions } from "./lib/queries/auth";
import { RouterProvider } from "@tanstack/react-router";
import { router } from "./router";

function InnerApp() {
  const { data: auth } = useSuspenseQuery(authQueryOptions);
  return <RouterProvider router={router} context={{ auth }} />;
}

export default function App() {
  return <InnerApp />;
}
