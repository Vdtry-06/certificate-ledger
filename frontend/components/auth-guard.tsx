"use client"

import type React from "react"

import { useEffect, useState } from "react"
import { useRouter, usePathname } from "next/navigation"
import { isLoggedIn } from "@/lib/auth"

interface AuthGuardProps {
  children: React.ReactNode
}

export function AuthGuard({ children }: AuthGuardProps) {
  const router = useRouter()
  const pathname = usePathname()
  const [checking, setChecking] = useState(true)

  useEffect(() => {
    const publicRoutes = ["/", "/login", "/register", "/verify"]

    const isPublicRoute = publicRoutes.includes(pathname) || pathname.startsWith("/verify")

    if (!isLoggedIn() && !isPublicRoute) {
      router.push("/login")
    }

    setChecking(false)
  }, [pathname, router])

  if (checking) {
    return null
  }

  return <>{children}</>
}
