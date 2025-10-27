import { queryOptions, mutationOptions } from "@tanstack/react-query";

export interface User {
  id: string;
  email: string;
  username: string;
}

export const authQueryOptions = queryOptions<User | null>({
  queryKey: ["auth"],
  queryFn: async () => null,
});

interface GoogleLoginResponse {
  redirectUrl: string;
}

export const googleLoginMutationOptions = mutationOptions({
  mutationFn: async (): Promise<GoogleLoginResponse> => {
    const response = await fetch(
      `${import.meta.env.VITE_BACKEND_URL}/google-login`,
      { method: "POST" }
    );

    if (!response.ok) {
      throw new Error("Failed to initiate Google login");
    }

    return response.json();
  },
});
