import { createFileRoute } from "@tanstack/react-router";
import { Hero } from "@/components/landing/Hero";
import { ProblemSection } from "@/components/landing/ProblemSection";
import { FeaturesSection } from "@/components/landing/FeaturesSection";
import { TargetUsersSection } from "@/components/landing/TargetUsersSection";
import { Footer } from "@/components/landing/Footer";

export const Route = createFileRoute("/")({
  component: LandingPage,
});

function LandingPage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100">
      <div className="max-w-6xl mx-auto px-6 py-16">
        <Hero />
        <ProblemSection />
        <FeaturesSection />
        <TargetUsersSection />
      </div>
      <Footer />
    </div>
  );
}
