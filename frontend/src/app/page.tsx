import Link from "next/link";

export default function Home() {
  return (
    <main className="flex flex-1 flex-col items-center justify-center gap-10 px-6 py-24 text-center">
      <p className="text-xs uppercase tracking-[0.4em] text-muted">
        a personal library ledger
      </p>
      <h1 className="heading-epic text-glow text-4xl text-neon sm:text-6xl">
        Who Has My Book
      </h1>
      <p className="cursor-blink max-w-md text-sm text-foreground/80">
        &gt; every book finds its way home
      </p>
      <Link
        href="/admin"
        className="border border-neon-dim px-6 py-2 text-sm uppercase tracking-[0.3em] text-neon transition hover:bg-neon hover:text-background"
      >
        Enter
      </Link>
    </main>
  );
}
