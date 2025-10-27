import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { User, Users, GraduationCap, Star, Zap, Shield } from "lucide-react";

export function FeaturesSection() {
  return (
    <div className="mb-20" id="features">
      <div className="text-center mb-12">
        <h2 className="text-4xl font-bold mb-4">How Trego Helps</h2>
        <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
          Powerful features designed for every member of the sports community
        </p>
      </div>

      <div className="grid md:grid-cols-2 gap-6">
        <Card className="group hover:shadow-lg transition-all duration-300 border-border/50 hover:border-primary/50">
          <CardHeader>
            <div className="flex items-start justify-between">
              <div className="size-12 rounded-xl bg-primary/10 flex items-center justify-center group-hover:bg-primary/20 transition-colors">
                <User className="size-6 text-primary" />
              </div>
              <Badge variant="secondary">
                <Zap className="size-3" />
                Smart Matching
              </Badge>
            </div>
            <CardTitle className="text-2xl mt-4">For Players</CardTitle>
          </CardHeader>
          <CardContent>
            <CardDescription className="text-base leading-relaxed">
              Create profiles showcasing your sports, skill levels, and
              availability. Find compatible teammates through smart matchmaking
              algorithms.
            </CardDescription>
          </CardContent>
        </Card>

        <Card className="group hover:shadow-lg transition-all duration-300 border-border/50 hover:border-primary/50">
          <CardHeader>
            <div className="flex items-start justify-between">
              <div className="size-12 rounded-xl bg-primary/10 flex items-center justify-center group-hover:bg-primary/20 transition-colors">
                <Users className="size-6 text-primary" />
              </div>
              <Badge variant="secondary">
                <Shield className="size-3" />
                Team Management
              </Badge>
            </div>
            <CardTitle className="text-2xl mt-4">For Teams & Clubs</CardTitle>
          </CardHeader>
          <CardContent>
            <CardDescription className="text-base leading-relaxed">
              Post openings, schedule tryouts, and manage substitutions all in
              one place. Build a reliable roster with accountability.
            </CardDescription>
          </CardContent>
        </Card>

        <Card className="group hover:shadow-lg transition-all duration-300 border-border/50 hover:border-primary/50">
          <CardHeader>
            <div className="flex items-start justify-between">
              <div className="size-12 rounded-xl bg-primary/10 flex items-center justify-center group-hover:bg-primary/20 transition-colors">
                <GraduationCap className="size-6 text-primary" />
              </div>
              <Badge variant="secondary">
                <Zap className="size-3" />
                Direct Booking
              </Badge>
            </div>
            <CardTitle className="text-2xl mt-4">For Coaches</CardTitle>
          </CardHeader>
          <CardContent>
            <CardDescription className="text-base leading-relaxed">
              Advertise your lessons and allow players to book sessions
              directly. Grow your network and find new opportunities.
            </CardDescription>
          </CardContent>
        </Card>

        <Card className="group hover:shadow-lg transition-all duration-300 border-border/50 hover:border-primary/50">
          <CardHeader>
            <div className="flex items-start justify-between">
              <div className="size-12 rounded-xl bg-primary/10 flex items-center justify-center group-hover:bg-primary/20 transition-colors">
                <Star className="size-6 text-primary" />
              </div>
              <Badge variant="secondary">
                <Shield className="size-3" />
                Trust & Safety
              </Badge>
            </div>
            <CardTitle className="text-2xl mt-4">Reputation System</CardTitle>
          </CardHeader>
          <CardContent>
            <CardDescription className="text-base leading-relaxed">
              Built-in reliability and reputation tracking ensures
              accountability and builds trust within the community.
            </CardDescription>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
