import type { Metadata } from "next";

export const metadata: Metadata = {
  title: "Admin",
};

export default function AdminPage() {
  return (
    <main className="mx-auto flex w-full max-w-4xl flex-1 flex-col gap-6 px-6 py-16">
      <h1 className="heading-epic text-glow text-3xl text-neon">Admin</h1>
      <p className="cursor-blink text-sm text-muted">
        &gt; books, friends and loans will live here
      </p>
    </main>
  );
}
