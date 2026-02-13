import { useKeenSlider } from 'keen-slider/react'

import { HomeContainer, Product } from '../styles/pages/home'

import 'keen-slider/keen-slider.min.css'
import Image from 'next/image'
import { stripe } from '@/lib/stripe'
import type { GetStaticProps } from 'next'
import type Stripe from 'stripe'
import Link from 'next/link'
import Head from 'next/head'
import { HandbagIcon } from '@phosphor-icons/react'
import { useShoppingCart } from 'use-shopping-cart'

interface HomeProps {
  products: {
    id: string
    name: string
    image: string
    price: number
    currency: 'BRL'
    formattedPrice: string
  }[]
}

export default function Home({ products }: Readonly<HomeProps>) {
  const { addItem } = useShoppingCart()
  const [sliderRef] = useKeenSlider({
    slides: {
      perView: 3,
      spacing: 48,
    },
  })

  return (
    <>
      <Head>
        <title>Home | Ignite Shop</title>
      </Head>

      <HomeContainer ref={sliderRef} className="keen-slider">
        {products.map(product => {
          return (
            <Link
              href={`/product/${product.id}`}
              key={product.id}
              prefetch={false}
            >
              <Product className="keen-slider__slide">
                <Image src={product.image} width={520} height={480} alt="" />

                <footer>
                  <div>
                    <strong>{product.name}</strong>
                    <span>{product.formattedPrice}</span>
                  </div>
                  <button
                    onClick={e => {
                      e.preventDefault()
                      addItem(product)
                    }}
                  >
                    <HandbagIcon size={32} weight="bold" />
                  </button>
                </footer>
              </Product>
            </Link>
          )
        })}
      </HomeContainer>
    </>
  )
}

export const getStaticProps: GetStaticProps = async () => {
  const response = await stripe.products.list({
    expand: ['data.default_price'],
  })

  const products = response.data.map(product => {
    const price = product.default_price as Stripe.Price

    return {
      id: product.id,
      name: product.name,
      image: product.images[0],
      price: price.unit_amount!,
      currency: 'BRL',
      formattedPrice: new Intl.NumberFormat('pt-BR', {
        style: 'currency',
        currency: 'BRL',
      }).format(price.unit_amount! / 100),
    }
  })

  return {
    props: {
      products,
    },
    revalidate: 60 * 60 * 2,
  }
}
