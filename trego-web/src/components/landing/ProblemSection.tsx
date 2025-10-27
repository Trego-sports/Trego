import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { AlertCircle, MessageSquareOff, UserX } from "lucide-react";

export function ProblemSection() {
  return (
    <div className="mb-20">
      <div className="text-center mb-12">
        <h2 className="text-4xl font-bold mb-4">The Problem We're Solving</h2>
        <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
          Sports enthusiasts face real challenges when trying to stay active and
          connected
        </p>
      </div>

      <div className="grid md:grid-cols-3 gap-6">
        <Card className="text-center border-border/50 hover:border-border transition-colors">
          <CardHeader>
            <div className="mx-auto mb-4 size-14 rounded-2xl bg-destructive/10 flex items-center justify-center">
              <UserX className="size-7 text-destructive" />
            </div>
            <CardTitle className="text-xl">Empty Courts</CardTitle>
          </CardHeader>
          <CardContent>
            <CardDescription className="text-base">
              Struggling to find reliable teammates to keep the game going
            </CardDescription>
          </CardContent>
        </Card>

        <Card className="text-center border-border/50 hover:border-border transition-colors">
          <CardHeader>
            <div className="mx-auto mb-4 size-14 rounded-2xl bg-warning/10 flex items-center justify-center">
              <MessageSquareOff className="size-7 text-warning-foreground" />
            </div>
            <CardTitle className="text-xl">Fragmented Info</CardTitle>
          </CardHeader>
          <CardContent>
            <CardDescription className="text-base">
              Tryout opportunities scattered across Instagram posts and group
              chats
            </CardDescription>
          </CardContent>
        </Card>

        <Card className="text-center border-border/50 hover:border-border transition-colors">
          <CardHeader>
            <div className="mx-auto mb-4 size-14 rounded-2xl bg-info/10 flex items-center justify-center">
              <AlertCircle className="size-7 text-info-foreground" />
            </div>
            <CardTitle className="text-xl">Lost Connections</CardTitle>
          </CardHeader>
          <CardContent>
            <CardDescription className="text-base">
              Difficulty maintaining consistent team rosters and finding coaches
            </CardDescription>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
