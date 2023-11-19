export function Footer() {
  return (
    <div className="px-16 py-8 w-full bg-secondary flex items-baseline justify-between">
      <div>
        <h1 className="font-bold text-lg">asest</h1>
      </div>

      <div>
        <a
          className="text-xs font-light hover:underline"
          target="_blank"
          href="https://www.caleblee.dev/"
        >
          Copyright © 2023 by Caleb Lee
        </a>
      </div>
    </div>
  );
}
