import type { Metadata } from "next";
import { Suspense } from "react";

export const metadata: Metadata = {
  title: "My borrowed books",
  // Friend pages are private links: keep them out of search engines.
  robots: { index: false, follow: false },
};

// The static shell (title/layout) is prerendered; the token-dependent part
// streams in inside <Suspense>.
export default function FriendPage(props: PageProps<"/friend/[token]">) {
  return (
    <main className="mx-auto flex w-full max-w-3xl flex-1 flex-col gap-6 px-6 py-16">
      <h1 className="heading-epic text-glow text-3xl text-neon">
        Books with you
      </h1>
      <Suspense
        fallback={<p className="cursor-blink text-sm text-muted">&gt; loading</p>}
      >
        <FriendLoans params={props.params} />
      </Suspense>
    </main>
  );
}

async function FriendLoans({
  params,
}: Pick<PageProps<"/friend/[token]">, "params">) {
  const { token } = await params;

  return (
    <p className="cursor-blink text-sm text-muted">
      &gt; loans for access key {token.slice(0, 6)}… coming soon
    </p>
  );
}
