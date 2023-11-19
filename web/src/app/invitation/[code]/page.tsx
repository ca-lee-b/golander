import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
} from "@/components/ui/card";

export default function Page({ params }: { params: { slug: string } }) {
  return (
    <div className="w-full min-h-screen flex flex-col items-center">
      <h1 className="text-4xl font-bold text-foreground my-8">
        You're invited!
      </h1>
      <Card className="border-border">
        <CardHeader className="flex flex-row gap-x-4">
          <div className="w-full p-2">
            <h1 className="text-2xl font-bold">Study Session at Bruinicks</h1>
            <p className="text-sm font-light">From Caleb Lee</p>
          </div>
          <img
            src="https://placehold.co/400"
            className="w-[100px] h-auto rounded-xl border-2 border-primary"
          />
        </CardHeader>

        <CardContent className="w-full">
          <p className="text-sm p-2">Choose your availability</p>
          <Calendar className="bg-accent w-max rounded-md" />
        </CardContent>

        <CardFooter className="w-full flex justify-center">
          <Button>Reserve</Button>
        </CardFooter>
      </Card>
    </div>
  );
}
