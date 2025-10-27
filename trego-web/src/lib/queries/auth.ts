import { queryOptions } from "@tanstack/react-query";

export interface User {
  id: string;
  email: string;
  username: string;
}

export const authQueryOptions = queryOptions<User | null>({
  queryKey: ["auth"],
  queryFn: async () => null,
});
