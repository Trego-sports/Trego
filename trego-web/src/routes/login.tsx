import { createFileRoute } from "@tanstack/react-router";
import { useMutation } from "@tanstack/react-query";
import { googleLoginMutationOptions } from "../lib/queries/auth";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Loader2 } from "lucide-react";

export const Route = createFileRoute("/login")({
  component: LoginPage,
});

function LoginPage() {
  const googleLoginMutation = useMutation(googleLoginMutationOptions);

  const handleGoogleSignIn = () => {
    googleLoginMutation.mutate(undefined, {
      onSuccess: (data) => {
        // Redirect to the Google OAuth URL
        window.location.href = data.redirectUrl;
      },
      onError: (error) => {
        console.error("Google login failed:", error);
        // TODO: Show error message to user
      },
    });
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-background via-muted/20 to-background p-4">
      <div className="w-full max-w-md">
        <div className="text-center mb-8 space-y-2">
          <h1 className="text-5xl font-bold bg-gradient-to-r from-primary to-primary/60 bg-clip-text text-transparent leading-tight">
            Trego
          </h1>
          <p className="text-muted-foreground text-lg">
            Your Hub for Sports Connections
          </p>
        </div>

        <Card className="border-border/50 shadow-xl">
          <CardHeader className="space-y-1">
            <CardTitle className="text-2xl text-center">Welcome back</CardTitle>
            <CardDescription className="text-center">
              Sign in to continue to your account
            </CardDescription>
          </CardHeader>
          <CardContent>
            <Button
              onClick={handleGoogleSignIn}
              disabled={googleLoginMutation.isPending}
              variant="outline"
              size="lg"
              className="w-full"
            >
              {googleLoginMutation.isPending ? (
                <>
                  <Loader2 className="animate-spin" />
                  Redirecting...
                </>
              ) : (
                <>
                  <img src="/google-logo.svg" alt="Google" className="size-5" />
                  Sign in with Google
                </>
              )}
            </Button>
          </CardContent>
        </Card>

        <div className="mt-8 text-center">
          <p className="text-sm text-muted-foreground">
            New to Trego?{" "}
            <a href="/" className="text-primary hover:underline font-medium">
              Learn more
            </a>
          </p>
        </div>
      </div>
    </div>
  );
}
