import { globalStyles } from '@/styles/global'
import type { AppProps } from 'next/app'
import logoImg from '../assets/logo.svg'
import { Container, Header, OpenCartButton } from '@/styles/pages/_app'
import Image from 'next/image'
import { useEffect, useState } from 'react'
import { CartProvider } from 'use-shopping-cart'
import { useRouter } from 'next/router'
import { HandbagIcon } from '@phosphor-icons/react'
import Cart from '@/components/Cart'

globalStyles()
export default function App({ Component, pageProps }: AppProps) {
  const [isCartOpen, setIsCartOpen] = useState(false)
  const [mounted, setMounted] = useState(false)
  const router = useRouter()

  useEffect(() => {
    // eslint-disable-next-line react-hooks/set-state-in-effect
    setMounted(true)
  }, [])

  return (
    <CartProvider
      mode="payment"
      cartMode="client-only"
      stripe={process.env.NEXT_PUBLIC_STRIPE_PUBLIC_KEY!}
      currency="BRL"
      successUrl={`${process.env.NEXT_PUBLIC_URL}/success`}
      cancelUrl={`${process.env.NEXT_PUBLIC_URL}/`}
      shouldPersist={true}
    >
      {mounted ? (
        <Container>
          <Header>
            <Image src={logoImg} alt="" onClick={() => router.push('/')} />

            <OpenCartButton onClick={() => setIsCartOpen(true)}>
              <HandbagIcon weight="bold" size={24} />
            </OpenCartButton>
          </Header>

          <Component {...pageProps} />
          {isCartOpen && <Cart onClose={() => setIsCartOpen(false)} />}
        </Container>
      ) : null}
    </CartProvider>
  )
}
