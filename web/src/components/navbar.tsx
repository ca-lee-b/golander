import Link from "next/link";
import { Button } from "./ui/button";
import { hasToken } from "@/lib/session";

export async function Navbar() {
  const token = await hasToken();

  return (
    <div className="py-4 px-32 flex w-full h-full items-baseline justify-around z-50 backdrop-blur-sm text-foreground drop-shadow-sm sticky mt-2 top-0 left-0">
      <div className="w-1/4">
        <Link href="/" className="font-bold text-md">
          asest
        </Link>
      </div>
      <div className="font-normal text-sm flex justify-between items-center w-1/2">
        <div className="flex items-center gap-x-1">
          <Link
            className="hover:bg-secondary py-2 px-3 rounded-sm transition duration-200"
            href="/"
          >
            Home
          </Link>
          <Link
            className="hover:bg-secondary py-2 px-3 rounded-sm transition duration-200"
            href="/"
          >
            Home
          </Link>
          <Link
            className="hover:bg-secondary py-2 px-3 rounded-sm transition duration-200"
            href="/"
          >
            Home
          </Link>
        </div>
        <div>
          {token ? (
            <div>
              <Button asChild variant="secondary">
                <Link href="/dashboard">Dashboard</Link>
              </Button>
            </div>
          ) : (
            <div>
              <Link
                className="hover:bg-secondary py-2 px-3 rounded-sm transition duration-200 mr-4"
                href="/login"
              >
                Login
              </Link>
              <Button className="text-sm duration-200" asChild size="sm">
                <Link href="/register">Register</Link>
              </Button>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
