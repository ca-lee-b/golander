import { cookies } from "next/headers";
export async function hasToken() {
  const cookie = cookies().get("token");

  if (!cookie || cookie.value.length < 1) {
    return false;
  }

  return true;
}
