"use client";

import { zodResolver } from "@hookform/resolvers/zod";
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
import { useRouter } from "next/navigation";
import Link from "next/link";

const registerSchema = z.object({
  email: z
    .string()
    .min(2)
    .max(50, {
      message: "Email cannot be longer than 50 characters",
    })
    .email({
      message: "Email must be an email",
    }),
  password: z
    .string()
    .min(7, {
      message: "Password must be at least 7 characters",
    })
    .max(50, {
      message: "Password cannot be longer than 50 characters",
    }),
});

export function RegisterForm() {
  const router = useRouter();
  const form = useForm<z.infer<typeof registerSchema>>({
    resolver: zodResolver(registerSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  function onSubmit(values: z.infer<typeof registerSchema>) {
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/register`, {
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
            {form.formState.errors.root.serverError.type === "404" &&
              "Email already exists"}
            {form.formState.errors.root.serverError.type === "500" &&
              "Internal server error! :("}
          </FormMessage>
        )}
        <Button className="w-full" type="submit">
          Register!
        </Button>
        <Link className="text-xs hover:underline" href="/login">
          Already have an account? Login
        </Link>
      </form>
    </Form>
  );
}
