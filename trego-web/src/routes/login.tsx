import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/login")({
  component: LoginPage,
});

function LoginPage() {
  const handleGoogleSignIn = () => {
    // TODO: Implement Google Sign-In logic
    console.log("Sign in with Google clicked");
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
            className="w-full flex items-center justify-center gap-3 px-6 py-3 bg-white border-2 border-slate-200 rounded-lg hover:border-slate-300 hover:shadow-md transition-all duration-200 group"
          >
            <img src="/google-logo.svg" alt="Google" className="w-5 h-5" />
            <span className="text-slate-700 font-medium group-hover:text-slate-900">
              Sign in with Google
            </span>
          </button>
        </div>
      </div>
    </div>
  );
}
