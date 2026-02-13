import { stripe } from '@/lib/stripe'
import {
  ImageContainer,
  ProductContainer,
  ProductDetails,
} from '@/styles/pages/product'
import type { GetStaticPaths, GetStaticProps } from 'next'
import Head from 'next/head'
import Image from 'next/image'
import type Stripe from 'stripe'
import { useShoppingCart } from 'use-shopping-cart'

interface ProductProps {
  product: {
    id: string
    name: string
    image: string
    price: number
    currency: 'BRL'
    formattedPrice: string
    description: string
    defaultPriceId: string
  }
}

export default function Product({ product }: ProductProps) {
  const { addItem } = useShoppingCart()

  return (
    <>
      <Head>
        <title>{product.name} | Ignite Shop</title>
      </Head>

      <ProductContainer>
        <ImageContainer>
          <Image src={product.image} alt="" width={520} height={480} />
        </ImageContainer>

        <ProductDetails>
          <h1>{product.name}</h1>
          <span>{product.formattedPrice}</span>
          <p>{product.description}</p>
          <button onClick={() => addItem(product)}>Colocar na sacola</button>
        </ProductDetails>
      </ProductContainer>
    </>
  )
}

export const getStaticPaths: GetStaticPaths = async () => {
  return {
    paths: [],
    fallback: true,
  }
}

export const getStaticProps: GetStaticProps = async ({ params }) => {
  const productId = params?.id as string
  const product = await stripe.products.retrieve(productId, {
    expand: ['default_price'],
  })

  const price = product.default_price as Stripe.Price

  return {
    props: {
      product: {
        id: product.id,
        name: product.name,
        image: product.images[0],
        price: price.unit_amount!,
        currency: 'BRL',
        formattedPrice: new Intl.NumberFormat('pt-BR', {
          style: 'currency',
          currency: 'BRL',
        }).format(price.unit_amount! / 100),
        description: product.description,
        defaultPriceId: price.id,
      },
    },
    revalidate: 60 * 60 * 1, // 1 hour
  }
}
