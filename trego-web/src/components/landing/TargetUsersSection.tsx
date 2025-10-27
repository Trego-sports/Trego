import { Link } from "@tanstack/react-router";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import {
  GraduationCap,
  Users,
  Trophy,
  Dumbbell,
  ArrowRight,
} from "lucide-react";

export function TargetUsersSection() {
  const userTypes = [
    { icon: GraduationCap, label: "University Students" },
    { icon: Dumbbell, label: "Recreational Players" },
    { icon: Trophy, label: "Sports Clubs" },
    { icon: Users, label: "Coaches" },
  ];

  return (
    <Card className="text-center border-border/50 bg-gradient-to-br from-card via-card to-muted/20">
      <CardContent className="py-12 px-8">
        <h2 className="text-4xl font-bold mb-4">Who We Serve</h2>
        <p className="text-lg text-muted-foreground max-w-3xl mx-auto mb-10">
          Trego is designed for university students, recreational players,
          sports clubs, and coaches seeking to grow their networks and find new
          opportunities to play.
        </p>

        <div className="flex flex-wrap justify-center gap-3 mb-10">
          {userTypes.map((type) => (
            <Badge
              key={type.label}
              variant="secondary"
              className="text-base px-4 py-2 gap-2"
            >
              <type.icon className="size-4" />
              {type.label}
            </Badge>
          ))}
        </div>

        <Button asChild size="lg" className="text-base">
          <Link to="/login">
            Join Trego Today
            <ArrowRight />
          </Link>
        </Button>
      </CardContent>
    </Card>
  );
}
