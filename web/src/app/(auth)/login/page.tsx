import { LoginForm } from "@/components/login-form";

export default function Page() {
  return (
    <div className=" text-foreground w-1/4 -mt-52">
      <h1 className="text-3xl px-4 font-bold mb-2 animate-in fade-in-0 slide-in-from-top-4 duration-500">
        Welcome back.
      </h1>
      <div className="bg-secondary w-full h-full rounded-md p-4">
        <h1 className="font-bold text-xl mb-4 rounded-md">Login</h1>
        <LoginForm />
      </div>
    </div>
  );
}
