import { authQueryOptions } from "@/lib/queries/auth";
import { useSuspenseQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/dashboard")({
  beforeLoad: async ({ context }) => {},
  component: DashboardPage,
});

function DashboardPage() {
  const { data: auth } = useSuspenseQuery(authQueryOptions);

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100">
      <div className="max-w-6xl mx-auto px-6 py-8">
        <div className="bg-white rounded-2xl shadow-lg p-8">
          <h1 className="text-3xl font-bold text-slate-900 mb-4">
            Welcome, {auth?.username}!
          </h1>
          <p className="text-slate-600">
            This is your dashboard. More features coming soon!
          </p>
        </div>
      </div>
    </div>
  );
}
