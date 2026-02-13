import { stripe } from '@/lib/stripe'
import { ImageContainer, SuccessContainer } from '@/styles/pages/success'
import type { GetServerSideProps } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import Link from 'next/link'
import { useEffect } from 'react'
import type Stripe from 'stripe'
import { useShoppingCart } from 'use-shopping-cart'

interface SuccessProps {
  customerName: string
  products: {
    id: string
    name: string
    imageUrl: string
    quantity: number
  }[]
}

export default function Success({
  customerName,
  products,
}: Readonly<SuccessProps>) {
  const { clearCart } = useShoppingCart()

  useEffect(() => {
    clearCart()
  }, [clearCart])

  const totalItemsCount = products.reduce(
    (total, product) => total + product.quantity,
    0,
  )

  return (
    <>
      <Head>
        <title>Compra efetuada | Ignite Shop</title>
        <meta name="robots" content="noindex" />
      </Head>

      <SuccessContainer>
        <div
          style={{
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            marginBottom: '3rem',
          }}
        >
          {products.map((product, index) => (
            <ImageContainer
              key={product.id}
              style={{
                marginLeft: index > 0 ? '-3.25rem' : 0,
                zIndex: products.length - index,
              }}
            >
              <Image src={product.imageUrl} alt="" width={120} height={110} />
            </ImageContainer>
          ))}
        </div>

        <h1>Compra efetuada</h1>

        <p>
          Uhuul <strong>{customerName}</strong>, sua compra de{' '}
          <strong>
            {totalItemsCount} {totalItemsCount > 1 ? 'camisetas' : 'camiseta'}
          </strong>{' '}
          já está a caminho de sua casa.
        </p>

        <Link href={'/'}>Voltar ao catálogo</Link>
      </SuccessContainer>
    </>
  )
}

export const getServerSideProps: GetServerSideProps = async ({ query }) => {
  if (!query.session_id) {
    return {
      redirect: { destination: '/', permanent: false },
    }
  }

  const sessionId = String(query.session_id)

  const session = await stripe.checkout.sessions.retrieve(sessionId, {
    expand: ['line_items', 'line_items.data.price.product'],
  })

  const customerName = session.customer_details?.name

  const products = session.line_items?.data.map(item => {
    const product = item.price?.product as Stripe.Product
    return {
      id: product.id,
      name: product.name,
      imageUrl: product.images[0],
      quantity: item.quantity,
    }
  })
  console.log(products)
  console.log(customerName)

  return {
    props: {
      customerName,
      products,
    },
  }
}
