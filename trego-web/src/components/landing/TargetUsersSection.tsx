import { Link } from "@tanstack/react-router";

export function TargetUsersSection() {
  return (
    <div className="bg-white rounded-2xl shadow-lg p-8 text-center">
      <h2 className="text-3xl font-bold text-slate-900 mb-6">Who We Serve</h2>
      <p className="text-lg text-slate-600 max-w-3xl mx-auto mb-8">
        Trego is designed for university students, recreational players, sports
        clubs, and coaches seeking to grow their networks and find new
        opportunities to play.
      </p>
      <Link
        to="/login"
        className="inline-block px-8 py-4 bg-slate-900 text-white font-semibold rounded-lg hover:bg-slate-800 transition-colors duration-200 shadow-lg hover:shadow-xl"
      >
        Join Trego Today
      </Link>
    </div>
  );
}
