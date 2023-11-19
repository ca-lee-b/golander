import { Button } from "@/components/ui/button";
import Link from "next/link";

export default function Home() {
  return (
    <main className="min-h-screen w-full">
      <article className="w-full h-screen bg-secondary py-2 px-16">
        <div className="w-full flex justify-between">
          <div className="text-7xl font-bold py-16 px-32 flex flex-col">
            <span className="animate-in fade-in-0 slide-in-from-left-4 duration-700 delay-0 fill-mode-both">
              A{" "}
            </span>
            <span className="text-primary animate-in fade-in-0 slide-in-from-left-4 duration-700 delay-200 fill-mode-both">
              Super-simple{" "}
            </span>
            <span className="animate-in fade-in-0 slide-in-from-left-4 duration-700 delay-500 fill-mode-both">
              Event{" "}
            </span>
            <span className="animate-in fade-in-0 slide-in-from-left-4 duration-700 delay-700 fill-mode-both">
              Scheduling{" "}
            </span>
            <span className="animate-in fade-in-0 slide-in-from-left-4 duration-700 delay-1000 fill-mode-both">
              Tool.
            </span>
          </div>
          <div className="animate-in fade-in-0 duration-700 delay-1250 py-16 px-32 rounded-md fill-mode-both">
            <img src="https://placehold.co/600x400" />
          </div>
        </div>

        <div className="w-full mt-40 animate-in fade-in-0 duration-700 delay-1250 py-16 px-32 rounded-md fill-mode-both">
          <div className="w-full flex items-center justify-center">
            <Button asChild className="" size="lg">
              <Link href="/register">Get Started</Link>
            </Button>
          </div>
        </div>
      </article>
    </main>
  );
}
