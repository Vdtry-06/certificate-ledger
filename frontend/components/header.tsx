"use client"

import Link from "next/link"
import { Award } from "lucide-react"
import { Button } from "@/components/ui/button"
import { isLoggedIn, logout } from "@/lib/auth"
import { usePathname } from "next/navigation"

export function Header() {
  const pathname = usePathname()
  const loggedIn = isLoggedIn()

  return (
    <header className="sticky top-0 z-10 border-b bg-background">
      <div className="container flex h-16 items-center justify-between py-4 px-4">
        <div className="flex items-center gap-2">
          <Award className="h-6 w-6" />
          <Link href="/">
            <h1 className="text-xl font-bold">Certificate Ledger</h1>
          </Link>
        </div>
        <nav className="flex items-center gap-4">
          {loggedIn && (
            <>
              <Link href="/dashboard">
                <Button variant={pathname === "/dashboard" ? "default" : "ghost"}>Dashboard</Button>
              </Link>
              <Link href="/issue">
                <Button variant={pathname === "/issue" ? "default" : "ghost"}>Issue Certificate</Button>
              </Link>
            </>
          )}
          <Link href="/verify">
            <Button variant={pathname === "/verify" ? "default" : "ghost"}>Verify</Button>
          </Link>
          {loggedIn ? (
            <Button onClick={logout}>Logout</Button>
          ) : (
            <Link href="/login">
              <Button>Login</Button>
            </Link>
          )}
        </nav>
      </div>
    </header>
  )
}