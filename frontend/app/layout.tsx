import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'
import { Providers } from '@/components/providers'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'GAuth - Authentication & Authorization Platform',
  description: 'Modern authentication and authorization management platform built on RFC-0150',
  keywords: ['authentication', 'authorization', 'gauth', 'security', 'identity'],
  authors: [{ name: 'Mauricio Fernandez' }],
  openGraph: {
    title: 'GAuth - Authentication & Authorization Platform',
    description: 'Modern authentication and authorization management platform',
    type: 'website',
  },
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className={inter.className}>
        <Providers>
          {children}
        </Providers>
      </body>
    </html>
  )
}