import { RegisterForm } from "@/components/register-form";
import { Card, CardHeader } from "@/components/ui/card";

export default function Page() {
  return (
    <div className=" text-foreground w-1/4 -mt-52">
      <h1 className="text-3xl px-4 font-bold mb-2 animate-in fade-in-0 slide-in-from-bottom-4 duration-500">
        Are you new here?
      </h1>
      <div className="bg-secondary w-full h-full rounded-md p-4">
        <h1 className="font-bold text-xl mb-4 rounded-md">Register</h1>
        <RegisterForm />
      </div>
    </div>
  );
}
