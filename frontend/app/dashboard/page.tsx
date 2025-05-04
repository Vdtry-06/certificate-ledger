"use client"

import { useEffect, useState } from "react"
import Link from "next/link"
import { Award, ChevronLeft, Download, Eye, Plus } from "lucide-react"
import { toast } from "sonner"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { getUserIssuedCertificates, getUserReceivedCertificates, type Certificate } from "@/lib/api"
import { getCurrentUser } from "@/lib/auth"
import { Header } from "@/components/header"
import { Footer } from "@/components/footer"

export default function Dashboard() {
  const [searchTerm, setSearchTerm] = useState("")
  const [isLoading, setIsLoading] = useState(true)
  const [issuedCertificates, setIssuedCertificates] = useState<Certificate[]>([])
  const [receivedCertificates, setReceivedCertificates] = useState<Certificate[]>([])

  useEffect(() => {
    const fetchCertificates = async () => {
      setIsLoading(true)
      try {

        const user = getCurrentUser()

        if (!user) {
          toast.error("Authentication Error", {
            description: "Please log in to view your dashboard.",
          })
          return
        }

        const issued = await getUserIssuedCertificates()
        setIssuedCertificates(issued)

        const received = await getUserReceivedCertificates(user.email)
        setReceivedCertificates(received)
      } catch (error) {
        toast.error("Error", {
          description: error instanceof Error ? error.message : "Failed to load certificates. Please try again.",
        })
      } finally {
        setIsLoading(false)
      }
    }

    fetchCertificates()
  }, [])

  const filteredIssued = issuedCertificates.filter(
    (cert) =>
      cert.recipientName?.toLowerCase().includes(searchTerm.toLowerCase()) ||
      cert.certificateTitle.toLowerCase().includes(searchTerm.toLowerCase()) ||
      cert.id.toLowerCase().includes(searchTerm.toLowerCase()),
  )

  const filteredReceived = receivedCertificates.filter(
    (cert) =>
      cert.issuerName?.toLowerCase().includes(searchTerm.toLowerCase()) ||
      cert.certificateTitle.toLowerCase().includes(searchTerm.toLowerCase()) ||
      cert.id.toLowerCase().includes(searchTerm.toLowerCase()),
  )

  return (
    <div className="flex min-h-screen flex-col">
    <Header />
    <div className="container max-w-5xl px-4 mx-auto py-10">
        <div className="mb-8">
          <Link href="/" className="flex items-center text-sm text-muted-foreground hover:text-foreground">
            <ChevronLeft className="mr-1 h-4 w-4" />
            Back to Home
          </Link>
        </div>
        <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between mb-8">
          <div className="flex items-center gap-2">
            <Award className="h-8 w-8" />
            <h1 className="text-3xl font-bold">Certificate Dashboard</h1>
          </div>
          <Link href="/issue">
            <Button>
              <Plus className="mr-2 h-4 w-4" />
              Issue New Certificate
            </Button>
          </Link>
        </div>
        <div className="mb-6">
          <Input
            placeholder="Search certificates..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="max-w-md"
          />
        </div>
        <Tabs defaultValue="issued">
          <TabsList className="mb-6">
            <TabsTrigger value="issued">Issued Certificates</TabsTrigger>
            <TabsTrigger value="received">Received Certificates</TabsTrigger>
          </TabsList>
          <TabsContent value="issued">
            <Card>
              <CardHeader>
                <CardTitle>Certificates You have Issued</CardTitle>
                <CardDescription>View and manage all certificates you have issued on the blockchain.</CardDescription>
              </CardHeader>
              <CardContent>
                {isLoading ? (
                  <div className="flex justify-center py-8">
                    <p>Loading certificates...</p>
                  </div>
                ) : filteredIssued.length > 0 ? (
                  <div className="rounded-md border">
                    <div className="grid grid-cols-5 gap-4 p-4 font-medium border-b">
                      <div>Certificate ID</div>
                      <div>Recipient</div>
                      <div>Title</div>
                      <div>Issue Date</div>
                      <div>Actions</div>
                    </div>
                    {filteredIssued.map((cert) => (
                      <div key={cert.id} className="grid grid-cols-5 gap-4 p-4 border-b last:border-0">
                        <div className="font-mono text-sm">{cert.id}</div>
                        <div>{cert.recipientName}</div>
                        <div>{cert.certificateTitle}</div>
                        <div>{new Date(cert.issueDate).toLocaleDateString()}</div>
                        <div className="flex gap-2">
                          <Link href={`/verify?id=${cert.id}`}>
                            <Button variant="outline" size="sm">
                              <Eye className="mr-1 h-3 w-3" />
                              View
                            </Button>
                          </Link>
                          <Button variant="outline" size="sm">
                            <Download className="mr-1 h-3 w-3" />
                            Export
                          </Button>
                        </div>
                      </div>
                    ))}
                  </div>
                ) : (
                  <div className="flex flex-col items-center justify-center py-12 text-center">
                    <Award className="h-12 w-12 text-muted-foreground mb-4" />
                    <h3 className="text-lg font-medium mb-2">No certificates found</h3>
                    <p className="text-muted-foreground mb-4">
                      {searchTerm
                        ? "No certificates match your search criteria."
                        : "You haven't issued any certificates yet."}
                    </p>
                    <Link href="/issue">
                      <Button>
                        <Plus className="mr-2 h-4 w-4" />
                        Issue New Certificate
                      </Button>
                    </Link>
                  </div>
                )}
              </CardContent>
            </Card>
          </TabsContent>
          <TabsContent value="received">
            <Card>
              <CardHeader>
                <CardTitle>Certificates You have Received</CardTitle>
                <CardDescription>View all certificates that have been issued to you.</CardDescription>
              </CardHeader>
              <CardContent>
                {isLoading ? (
                  <div className="flex justify-center py-8">
                    <p>Loading certificates...</p>
                  </div>
                ) : filteredReceived.length > 0 ? (
                  <div className="rounded-md border">
                    <div className="grid grid-cols-5 gap-4 p-4 font-medium border-b">
                      <div>Certificate ID</div>
                      <div>Issuer</div>
                      <div>Title</div>
                      <div>Issue Date</div>
                      <div>Actions</div>
                    </div>
                    {filteredReceived.map((cert) => (
                      <div key={cert.id} className="grid grid-cols-5 gap-4 p-4 border-b last:border-0">
                        <div className="font-mono text-sm">{cert.id}</div>
                        <div>{cert.issuerName}</div>
                        <div>{cert.certificateTitle}</div>
                        <div>{new Date(cert.issueDate).toLocaleDateString()}</div>
                        <div className="flex gap-2">
                          <Link href={`/verify?id=${cert.id}`}>
                            <Button variant="outline" size="sm">
                              <Eye className="mr-1 h-3 w-3" />
                              View
                            </Button>
                          </Link>
                          <Button variant="outline" size="sm">
                            <Download className="mr-1 h-3 w-3" />
                            Export
                          </Button>
                        </div>
                      </div>
                    ))}
                  </div>
                ) : (
                  <div className="flex flex-col items-center justify-center py-12 text-center">
                    <Award className="h-12 w-12 text-muted-foreground mb-4" />
                    <h3 className="text-lg font-medium mb-2">No certificates found</h3>
                    <p className="text-muted-foreground">
                      {searchTerm
                        ? "No certificates match your search criteria."
                        : "You haven't received any certificates yet."}
                    </p>
                  </div>
                )}
              </CardContent>
            </Card>
          </TabsContent>
        </Tabs>
      </div>
      <Footer />
    </div>
  )
}
