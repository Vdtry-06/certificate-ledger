"use client"

import type React from "react"

import { useState } from "react"
import Link from "next/link"
import { ChevronLeft, FileCheck, Search } from "lucide-react"
import { toast } from "sonner"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { getCertificate, verifyCertificate, type Certificate } from "@/lib/api"
import { Header } from "@/components/header"
import { useSearchParams } from "next/navigation"
import { useEffect } from "react"

export default function VerifyCertificate() {
  const searchParams = useSearchParams()
  const [certificateId, setCertificateId] = useState("")
  const [isVerifying, setIsVerifying] = useState(false)
  const [certificate, setCertificate] = useState<Certificate | null>(null)
  const [isValid, setIsValid] = useState(false)

  useEffect(() => {
    const id = searchParams.get("id")
    if (id) {
      setCertificateId(id)
      handleVerify(null, id)
    }
  }, [searchParams])

  const handleVerify = async (e: React.FormEvent | null, idOverride?: string) => {
    if (e) e.preventDefault()

    const idToVerify = idOverride || certificateId
    if (!idToVerify.trim()) return

    setIsVerifying(true)
    setCertificate(null)
    setIsValid(false)

    try {
      try {
        const cert = await getCertificate(idToVerify)
        setCertificate(cert)

        const verification = await verifyCertificate(cert.hash)
        setIsValid(verification.valid)

        if (verification.valid) {
          toast.success("Certificate Verified", {
            description: "This certificate has been verified on the blockchain.",
          })
        } else {
          toast.error("Certificate Invalid", {
            description: "This certificate could not be verified on the blockchain.",
          })
        }
      } catch {
        if (idToVerify.startsWith("0x")) {
          try {
            const verification = await verifyCertificate(idToVerify)
            if (verification.valid) {
              setIsValid(true)
              toast.success("Certificate is valid", {
                description: "This certificate has been verified on the blockchain.",
              })
            } else {
              toast.error("Invalid Certificate", {
                description: "This certificate could not be verified on the blockchain.",
              })
            }
          } catch {
            toast.error("Verification Error", {
              description: "Failed to verify certificate hash. Please try again.",
            })
          }
        } else {
          toast.error("Certificate Not Found", {
            description: "No certificate found with the provided ID or hash.",
          })
        }
      }
    } catch (error) {
      toast.error("Verification Error", {
        description: error instanceof Error ? error.message : "Failed to verify certificate. Please try again.",
      })
    } finally {
      setIsVerifying(false)
    }
  }

  return (
    <div className="flex min-h-screen flex-col">
      <Header />
      <div className="container max-w-4xl px-4 mx-auto py-10">
        <div className="mb-8">
          <Link href="/" className="flex items-center text-sm text-muted-foreground hover:text-foreground">
            <ChevronLeft className="mr-1 h-4 w-4" />
            Back to Home
          </Link>
        </div>
        <div className="flex items-center gap-2 mb-8">
          <FileCheck className="h-8 w-8" />
          <h1 className="text-3xl font-bold">Verify Certificate</h1>
        </div>
        <Card className="mb-8">
          <CardHeader>
            <CardTitle>Certificate Verification</CardTitle>
            <CardDescription>
              Enter a certificate ID or hash to verify its authenticity on the blockchain.
            </CardDescription>
          </CardHeader>
          <CardContent>
            <form onSubmit={(e) => handleVerify(e)} className="flex gap-2">
              <Input
                placeholder="Enter certificate ID or hash"
                value={certificateId}
                onChange={(e) => setCertificateId(e.target.value)}
                className="flex-1"
              />
              <Button type="submit" disabled={isVerifying}>
                {isVerifying ? (
                  "Verifying..."
                ) : (
                  <>
                    <Search className="mr-2 h-4 w-4" />
                    Verify
                  </>
                )}
              </Button>
            </form>
          </CardContent>
        </Card>

        {certificate && (
          <Card>
            <CardHeader className="border-b">
              <div className="flex items-center justify-between">
                <CardTitle>Certificate Details</CardTitle>
                <div
                  className={`flex items-center gap-2 rounded-full px-3 py-1 text-sm ${
                    isValid ? "bg-green-100 text-green-800" : "bg-red-100 text-red-800"
                  }`}
                >
                  <FileCheck className="h-4 w-4" />
                  {isValid ? "Verified" : "Not Verified"}
                </div>
              </div>
            </CardHeader>
            <CardContent className="pt-6">
              <div className="grid gap-4 md:grid-cols-2">
                <div className="space-y-1">
                  <p className="text-sm font-medium text-muted-foreground">Certificate ID</p>
                  <p>{certificate.id}</p>
                </div>
                <div className="space-y-1">
                  <p className="text-sm font-medium text-muted-foreground">Blockchain Hash</p>
                  <p className="truncate">{certificate.hash}</p>
                </div>
                <div className="space-y-1">
                  <p className="text-sm font-medium text-muted-foreground">Recipient</p>
                  <p>{certificate.recipientName}</p>
                </div>
                <div className="space-y-1">
                  <p className="text-sm font-medium text-muted-foreground">Recipient Email</p>
                  <p>{certificate.recipientEmail}</p>
                </div>
                <div className="space-y-1">
                  <p className="text-sm font-medium text-muted-foreground">Certificate Title</p>
                  <p>{certificate.certificateTitle}</p>
                </div>
                <div className="space-y-1">
                  <p className="text-sm font-medium text-muted-foreground">Issue Date</p>
                  <p>{new Date(certificate.issueDate).toLocaleDateString()}</p>
                </div>
                <div className="space-y-1">
                  <p className="text-sm font-medium text-muted-foreground">Issuer</p>
                  <p>{certificate.issuerName}</p>
                </div>
                <div className="space-y-1">
                  <p className="text-sm font-medium text-muted-foreground">Block Number</p>
                  <p>{certificate.blockNumber}</p>
                </div>
              </div>
              <div className="mt-6 space-y-1">
                <p className="text-sm font-medium text-muted-foreground">Description</p>
                <p>{certificate.description}</p>
              </div>
              <div className="mt-6 rounded-lg bg-muted p-4">
                <p className="text-sm">
                  {isValid
                    ? "This certificate has been verified on the blockchain. The certificate data is immutable and cannot be altered."
                    : "This certificate could not be verified on the blockchain. It may have been tampered with or does not exist."}
                </p>
              </div>
            </CardContent>
          </Card>
        )}
      </div>
    </div>
  )
}
