import { useEffect } from "react";
import type { AppProps } from "next/app";
import "~/styles/globals.css";
import {
  ClerkProvider,
  SignedIn,
  SignedOut,
  SignInButton,
  UserButton,
  useAuth,
} from "@clerk/nextjs";

function Header() {
  const { getToken } = useAuth();
  useEffect(() => {
    const fetch = async () => {
      const token = await getToken({ template: "jwt" });
      console.log("🚀 ~ file: _app.tsx:18 ~ fetch ~ token:", token);
    };

    fetch();
  }, []);

  return (
    <header
      style={{ display: "flex", justifyContent: "space-between", padding: 20 }}
    >
      <h1>Supertruyen</h1>
      <SignedIn>
        {/* Mount the UserButton component */}
        <UserButton />
      </SignedIn>
      <SignedOut>
        {/* Signed out users get sign in button */}
        <SignInButton />
      </SignedOut>
    </header>
  );
}

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ClerkProvider {...pageProps}>
      <Header />
      <Component {...pageProps} />
    </ClerkProvider>
  );
}

export default MyApp;
