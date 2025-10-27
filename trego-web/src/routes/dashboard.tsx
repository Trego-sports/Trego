import { authQueryOptions } from "@/lib/queries/auth";
import { useSuspenseQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import {
  User,
  Calendar,
  Users,
  Trophy,
  Settings,
  Bell,
  Activity,
  TrendingUp,
} from "lucide-react";

export const Route = createFileRoute("/dashboard")({
  beforeLoad: async ({}) => {},
  component: DashboardPage,
});

function DashboardPage() {
  const { data: auth } = useSuspenseQuery(authQueryOptions);

  return (
    <div className="min-h-screen bg-gradient-to-br from-background via-muted/20 to-background">
      {/* Header */}
      <header className="border-b bg-card/50 backdrop-blur-sm sticky top-0 z-10">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-4">
              <h1 className="text-2xl font-bold">Trego</h1>
              <Badge variant="secondary">Beta</Badge>
            </div>
            <div className="flex items-center gap-2">
              <Button variant="ghost" size="icon">
                <Bell className="size-5" />
              </Button>
              <Button variant="ghost" size="icon">
                <Settings className="size-5" />
              </Button>
              <div className="flex items-center gap-2 ml-2">
                <div className="size-8 rounded-full bg-primary/10 flex items-center justify-center">
                  <User className="size-4 text-primary" />
                </div>
                <span className="text-sm font-medium">{auth?.username}</span>
              </div>
            </div>
          </div>
        </div>
      </header>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Welcome Section */}
        <div className="mb-8">
          <h2 className="text-3xl font-bold mb-2">
            Welcome back, {auth?.username}! 👋
          </h2>
          <p className="text-muted-foreground">
            Here's what's happening with your sports connections today.
          </p>
        </div>

        {/* Stats Grid */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
          <Card>
            <CardHeader>
              <div className="flex items-center justify-between">
                <CardTitle className="text-sm font-medium">
                  Active Games
                </CardTitle>
                <Calendar className="size-4 text-muted-foreground" />
              </div>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">12</div>
              <p className="text-xs text-muted-foreground mt-1">
                <span className="text-success-foreground">+2</span> from last
                week
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <div className="flex items-center justify-between">
                <CardTitle className="text-sm font-medium">
                  Connections
                </CardTitle>
                <Users className="size-4 text-muted-foreground" />
              </div>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">48</div>
              <p className="text-xs text-muted-foreground mt-1">
                <span className="text-success-foreground">+5</span> new this
                month
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <div className="flex items-center justify-between">
                <CardTitle className="text-sm font-medium">
                  Reputation Score
                </CardTitle>
                <Trophy className="size-4 text-muted-foreground" />
              </div>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">4.8</div>
              <p className="text-xs text-muted-foreground mt-1">
                Based on 24 reviews
              </p>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <div className="flex items-center justify-between">
                <CardTitle className="text-sm font-medium">Activity</CardTitle>
                <Activity className="size-4 text-muted-foreground" />
              </div>
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">89%</div>
              <p className="text-xs text-muted-foreground mt-1">
                <span className="text-success-foreground">+12%</span> this month
              </p>
            </CardContent>
          </Card>
        </div>

        {/* Main Content Grid */}
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
          {/* Upcoming Games */}
          <Card className="lg:col-span-2">
            <CardHeader>
              <CardTitle>Upcoming Games</CardTitle>
              <CardDescription>
                Your scheduled matches and practice sessions
              </CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                <div className="flex items-center justify-between p-4 rounded-lg border bg-muted/30">
                  <div className="flex items-center gap-4">
                    <div className="size-10 rounded-lg bg-primary/10 flex items-center justify-center">
                      <Trophy className="size-5 text-primary" />
                    </div>
                    <div>
                      <p className="font-medium">Basketball League Match</p>
                      <p className="text-sm text-muted-foreground">
                        Tomorrow at 6:00 PM
                      </p>
                    </div>
                  </div>
                  <Button size="sm">View Details</Button>
                </div>

                <div className="flex items-center justify-between p-4 rounded-lg border bg-muted/30">
                  <div className="flex items-center gap-4">
                    <div className="size-10 rounded-lg bg-primary/10 flex items-center justify-center">
                      <Users className="size-5 text-primary" />
                    </div>
                    <div>
                      <p className="font-medium">Soccer Practice</p>
                      <p className="text-sm text-muted-foreground">
                        Friday at 4:30 PM
                      </p>
                    </div>
                  </div>
                  <Button size="sm">View Details</Button>
                </div>

                <div className="flex items-center justify-between p-4 rounded-lg border bg-muted/30">
                  <div className="flex items-center gap-4">
                    <div className="size-10 rounded-lg bg-primary/10 flex items-center justify-center">
                      <Calendar className="size-5 text-primary" />
                    </div>
                    <div>
                      <p className="font-medium">Tennis Doubles</p>
                      <p className="text-sm text-muted-foreground">
                        Saturday at 10:00 AM
                      </p>
                    </div>
                  </div>
                  <Button size="sm">View Details</Button>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Quick Actions */}
          <Card>
            <CardHeader>
              <CardTitle>Quick Actions</CardTitle>
              <CardDescription>Get started with common tasks</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-2">
                <Button variant="outline" className="w-full justify-start">
                  <Calendar className="size-4" />
                  Schedule a Game
                </Button>
                <Button variant="outline" className="w-full justify-start">
                  <Users className="size-4" />
                  Find Teammates
                </Button>
                <Button variant="outline" className="w-full justify-start">
                  <Trophy className="size-4" />
                  Join a Tournament
                </Button>
                <Button variant="outline" className="w-full justify-start">
                  <User className="size-4" />
                  Update Profile
                </Button>
              </div>
            </CardContent>
          </Card>
        </div>

        {/* Coming Soon Banner */}
        <Card className="mt-6 bg-gradient-to-r from-primary/10 via-primary/5 to-primary/10 border-primary/20">
          <CardContent className="flex items-center justify-between py-4">
            <div className="flex items-center gap-4">
              <TrendingUp className="size-8 text-primary" />
              <div>
                <p className="font-semibold">More features coming soon!</p>
                <p className="text-sm text-muted-foreground">
                  We're working hard to bring you advanced matchmaking, team
                  management, and more.
                </p>
              </div>
            </div>
            <Button>Learn More</Button>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
