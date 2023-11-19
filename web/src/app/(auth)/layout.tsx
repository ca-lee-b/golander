import { hasToken } from "@/lib/session";
import { redirect } from "next/navigation";
import React from "react";

export default async function Layout({
  children,
}: {
  children: React.ReactNode;
}) {
  const token = await hasToken();

  if (token) {
    redirect("/");
  }

  return (
    <div className="w-full h-screen flex items-center justify-center ease-out">
      {children}
    </div>
  );
}
