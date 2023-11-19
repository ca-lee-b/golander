import { DashboardSidebar } from "@/components/dashboard-sidebar";
import { hasToken } from "@/lib/session";
import { redirect } from "next/navigation";

export default async function Page() {
  const token = await hasToken();

  if (!token) {
    redirect("/login");
  }

  return (
    <div className="w-full min-h-screen bg-background px-32 py-12">
      <h1 className="text-4xl font-bold mb-8">Welcome back, Caleb</h1>

      <DashboardSidebar />
    </div>
  );
}
