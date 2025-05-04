"use client"

import Link from "next/link"
import { Award } from "lucide-react"

export function Footer() {
  return (
    <footer className="border-t py-6 md:py-8">
      <div className="container flex flex-col items-center justify-between gap-4 md:flex-row px-4">
        <div className="flex items-center gap-2">
          <Award className="h-5 w-5" />
          <p className="text-sm text-muted-foreground">© 2025 Certificate Ledger. All rights reserved.</p>
        </div>
        <nav className="flex gap-4">
          <Link href="#" className="text-sm text-muted-foreground hover:underline">
            Terms
          </Link>
          <Link href="#" className="text-sm text-muted-foreground hover:underline">
            Privacy
          </Link>
          <Link href="#" className="text-sm text-muted-foreground hover:underline">
            Contact
          </Link>
        </nav>
      </div>
    </footer>
  )
}