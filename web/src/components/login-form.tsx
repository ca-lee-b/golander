"use client";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import * as z from "zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "./ui/form";
import { Input } from "./ui/input";
import { Button } from "./ui/button";
import Link from "next/link";

const loginSchema = z.object({
  email: z
    .string()
    .min(2, {
      message: "Email is required",
    })
    .email({
      message: "Email is required",
    }),
  password: z.string().min(1, {
    message: "Password is required",
  }),
});

export function LoginForm() {
  const router = useRouter();
  const form = useForm<z.infer<typeof loginSchema>>({
    resolver: zodResolver(loginSchema),
  });

  function onSubmit(values: z.infer<typeof loginSchema>) {
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/login`, {
      method: "POST",
      credentials: "include",
      cache: "no-cache",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: values.email,
        password: values.password,
      }),
    }).then(async (res) => {
      if (res.status > 200) {
        const status = await res.text();
        form.setError("root.serverError", {
          type: res.status.toString(),
        });
        return;
      }

      router.refresh();
      router.push("/");
    });
  }

  return (
    <Form {...form}>
      <form className="space-y-10" onSubmit={form.handleSubmit(onSubmit)}>
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel className="font-normal">Email</FormLabel>
              <FormControl>
                <Input placeholder="example@email.com" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel className="font-normal">Password</FormLabel>
              <FormControl>
                <Input type="password" placeholder="shhh!" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        {form.formState.errors.root && (
          <FormMessage>
            {form.formState.errors.root.type === "404" &&
              "Email does not exist"}
          </FormMessage>
        )}
        <Button className="w-full" type="submit">
          Login
        </Button>
        <Link className="text-xs hover:underline" href="/register">
          Don't have an account? Register here
        </Link>
      </form>
    </Form>
  );
}
