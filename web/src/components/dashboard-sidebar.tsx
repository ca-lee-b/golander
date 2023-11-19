import Link from "next/link";
import { Button } from "./ui/button";
import { HomeIcon } from "@radix-ui/react-icons";

export function DashboardSidebar() {
  return (
    <div className="h-max w-1/6 bg-secondary rounded-md">
      <div className="flex flex-col items-center justify-center">
        <Button
          variant="secondary"
          size="lg"
          className="w-full h-auto py-4 rounded-none rounded-tr-md rounded-tl-md justify-start shadow-none gap-x-5 hover:bg-background/40 duration-150"
        >
          <HomeIcon className="w-[20px] h-[20px]" />
          Home
        </Button>
        <Button
          variant="secondary"
          size="lg"
          className="w-full h-auto py-4 rounded-none justify-start shadow-none gap-x-5 hover:bg-background/40 duration-150"
        >
          <HomeIcon className="w-[20px] h-[20px]" />
          Home
        </Button>
        <Button
          variant="secondary"
          size="lg"
          className="w-full h-auto py-4 rounded-none justify-start shadow-none gap-x-5 hover:bg-background/40 duration-150"
        >
          <HomeIcon className="w-[20px] h-[20px]" />
          Home
        </Button>
        <Button
          variant="secondary"
          size="lg"
          className="w-full h-auto py-4 rounded-none justify-start shadow-none gap-x-5 hover:bg-background/40 duration-150"
        >
          <HomeIcon className="w-[20px] h-[20px]" />
          Home
        </Button>
        <Button
          variant="secondary"
          size="lg"
          className="w-full h-auto py-4 rounded-none justify-start shadow-none gap-x-5 hover:bg-background/40 duration-150"
        >
          <HomeIcon className="w-[20px] h-[20px]" />
          Home
        </Button>
        <Button
          variant="secondary"
          size="lg"
          className="w-full h-auto py-4 rounded-none justify-start shadow-none gap-x-5 hover:bg-background/40 duration-150"
        >
          <HomeIcon className="w-[20px] h-[20px]" />
          Home
        </Button>
        <Button
          variant="secondary"
          size="lg"
          className="w-full h-auto py-4 rounded-none justify-start shadow-none gap-x-5 hover:bg-background/40 duration-150"
        >
          <HomeIcon className="w-[20px] h-[20px]" />
          Home
        </Button>
      </div>
    </div>
  );
}
