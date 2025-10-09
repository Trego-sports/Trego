export function ProblemSection() {
  return (
    <div className="mb-16 bg-white rounded-2xl shadow-lg p-8">
      <h2 className="text-3xl font-bold text-slate-900 mb-6 text-center">
        The Problem
      </h2>
      <div className="grid md:grid-cols-3 gap-6">
        <div className="text-center">
          <div className="text-4xl mb-4">🏀</div>
          <h3 className="font-semibold text-lg text-slate-900 mb-2">
            Empty Courts
          </h3>
          <p className="text-slate-600">
            Struggling to find reliable teammates to keep the game going
          </p>
        </div>
        <div className="text-center">
          <div className="text-4xl mb-4">📱</div>
          <h3 className="font-semibold text-lg text-slate-900 mb-2">
            Fragmented Info
          </h3>
          <p className="text-slate-600">
            Tryout opportunities scattered across Instagram posts and group
            chats
          </p>
        </div>
        <div className="text-center">
          <div className="text-4xl mb-4">🤝</div>
          <h3 className="font-semibold text-lg text-slate-900 mb-2">
            Lost Connections
          </h3>
          <p className="text-slate-600">
            Difficulty maintaining consistent team rosters and finding coaches
          </p>
        </div>
      </div>
    </div>
  );
}
