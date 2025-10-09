import { Link } from "@tanstack/react-router";

export function Hero() {
  return (
    <div className="text-center mb-16">
      <h1 className="text-6xl font-bold text-slate-900 mb-4">Trego</h1>
      <p className="text-2xl text-slate-600 mb-8">
        Your Hub for Sports Connections
      </p>
      <p className="text-lg text-slate-500 max-w-3xl mx-auto mb-12">
        Connect with players, teams, clubs, and coaches all in one place. Find
        reliable teammates, discover tryout opportunities, and build lasting
        sports connections.
      </p>
      <Link
        to="/login"
        className="inline-block px-8 py-4 bg-slate-900 text-white font-semibold rounded-lg hover:bg-slate-800 transition-colors duration-200 shadow-lg hover:shadow-xl"
      >
        Get Started
      </Link>
    </div>
  );
}
