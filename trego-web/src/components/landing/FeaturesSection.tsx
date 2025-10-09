export function FeaturesSection() {
  return (
    <div className="mb-16">
      <h2 className="text-3xl font-bold text-slate-900 mb-8 text-center">
        How Trego Helps
      </h2>
      <div className="grid md:grid-cols-2 gap-6">
        <div className="bg-white rounded-xl shadow-md p-6 hover:shadow-lg transition-shadow">
          <h3 className="text-xl font-semibold text-slate-900 mb-3">
            For Players
          </h3>
          <p className="text-slate-600">
            Create profiles showcasing your sports, skill levels, and
            availability. Find compatible teammates through smart matchmaking
            algorithms.
          </p>
        </div>
        <div className="bg-white rounded-xl shadow-md p-6 hover:shadow-lg transition-shadow">
          <h3 className="text-xl font-semibold text-slate-900 mb-3">
            For Teams & Clubs
          </h3>
          <p className="text-slate-600">
            Post openings, schedule tryouts, and manage substitutions all in one
            place. Build a reliable roster with accountability.
          </p>
        </div>
        <div className="bg-white rounded-xl shadow-md p-6 hover:shadow-lg transition-shadow">
          <h3 className="text-xl font-semibold text-slate-900 mb-3">
            For Coaches
          </h3>
          <p className="text-slate-600">
            Advertise your lessons and allow players to book sessions directly.
            Grow your network and find new opportunities.
          </p>
        </div>
        <div className="bg-white rounded-xl shadow-md p-6 hover:shadow-lg transition-shadow">
          <h3 className="text-xl font-semibold text-slate-900 mb-3">
            Reputation System
          </h3>
          <p className="text-slate-600">
            Built-in reliability and reputation tracking ensures accountability
            and builds trust within the community.
          </p>
        </div>
      </div>
    </div>
  );
}
