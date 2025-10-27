import { Link } from "@tanstack/react-router";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { ArrowRight, Sparkles } from "lucide-react";

export function Hero() {
  return (
    <div className="text-center mb-20">

      <Badge variant="secondary" className="mb-6 gap-1.5">
        <Sparkles className="size-3" />
        Now in Beta
      </Badge>

      <h1 className="text-6xl sm:text-7xl lg:text-8xl font-bold mb-6 text-foreground">
        Trego
      </h1>

      <p className="text-2xl sm:text-3xl lg:text-4xl font-semibold text-foreground/90 mb-6">
        Your Hub for Sports Connections
      </p>

      <p className="text-lg sm:text-xl text-muted-foreground max-w-3xl mx-auto mb-12 leading-relaxed">
        Connect with players, teams, clubs, and coaches all in one place. Find
        reliable teammates, discover tryout opportunities, and build lasting
        sports connections.
      </p>

      <div className="flex flex-col sm:flex-row gap-4 justify-center items-center">
        <Button asChild size="lg" className="text-base">
          <Link to="/login">
            Get Started
            <ArrowRight />
          </Link>
        </Button>
        <Button asChild variant="outline" size="lg" className="text-base">
          <a href="#features">Learn More</a>
        </Button>
      </div>
    </div>
  );
}
