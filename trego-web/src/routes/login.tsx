import { createFileRoute } from "@tanstack/react-router";
import { useMutation } from "@tanstack/react-query";
import { googleLoginMutationOptions } from "../lib/queries/auth";

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
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-slate-50 to-slate-100">
      <div className="w-full max-w-md px-8">
        <div className="text-center mb-8">
          <h1 className="text-4xl font-bold text-slate-900 mb-2">Trego</h1>
          <p className="text-slate-600">Sign in to continue</p>
        </div>

        <div className="bg-white rounded-2xl shadow-lg p-8">
          <button
            onClick={handleGoogleSignIn}
            disabled={googleLoginMutation.isPending}
            className="w-full flex items-center justify-center gap-3 px-6 py-3 bg-white border-2 border-slate-200 rounded-lg hover:border-slate-300 hover:shadow-md transition-all duration-200 group disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <img src="/google-logo.svg" alt="Google" className="w-5 h-5" />
            <span className="text-slate-700 font-medium group-hover:text-slate-900">
              {googleLoginMutation.isPending
                ? "Redirecting..."
                : "Sign in with Google"}
            </span>
          </button>
        </div>
      </div>
    </div>
  );
}
