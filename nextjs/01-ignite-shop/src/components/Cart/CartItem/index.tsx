import Image from 'next/image'
import {
  CartItemContainer,
  ImageContainer,
  ProductInfo,
  ProductControls,
} from './styles'
import { useShoppingCart } from 'use-shopping-cart'

export interface CartItemProps {
  product: {
    id: string
    name: string
    image: string
    price: number
    quantity: number
    currency: 'BRL'
    formattedPrice: string
  }
}

export default function CartItem({ product }: Readonly<CartItemProps>) {
  const { removeItem, decrementItem, incrementItem } = useShoppingCart()

  return (
    <CartItemContainer>
      <ImageContainer>
        <Image src={product.image} alt="" width={95} height={95} />
      </ImageContainer>
      <ProductInfo>
        <p>{product.name}</p>
        <span>{product.formattedPrice}</span>
        <ProductControls>
          <button onClick={() => removeItem(product.id)}>Remover</button>
          <div>
            <button onClick={() => decrementItem(product.id)}>-</button>
            <span>{product.quantity}</span>
            <button onClick={() => incrementItem(product.id)}>+</button>
          </div>
        </ProductControls>
      </ProductInfo>
    </CartItemContainer>
  )
}
