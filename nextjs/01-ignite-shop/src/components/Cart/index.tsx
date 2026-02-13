import { XIcon } from '@phosphor-icons/react'
import {
  CartContainer,
  CloseCartButton,
  Footer,
  ItemsListContainer,
  PurchaseButton,
  TotalContainer,
} from './styles'
import CartItem from './CartItem'
import { useShoppingCart } from 'use-shopping-cart'
import { useState } from 'react'
import axios from 'axios'

interface CartProps {
  onClose: () => void
}

export default function Cart({ onClose }: Readonly<CartProps>) {
  const { cartCount, cartDetails, formattedTotalPrice } = useShoppingCart()
  const cartItems = Object.values(cartDetails ?? {})

  const [isCreatingCheckoutSession, setIsCreatingCheckoutSession] =
    useState(false)

  async function handleBuyProduct() {
    setIsCreatingCheckoutSession(true)
    try {
      const response = await axios.post('/api/checkout', {
        cartDetails,
      })

      const { checkoutUrl } = response.data

      globalThis.location.href = checkoutUrl
    } catch {
      setIsCreatingCheckoutSession(false)
      alert('falha ao redirecionar ao checkout!')
    }
  }

  return (
    <CartContainer>
      <CloseCartButton onClick={onClose}>
        <XIcon size={24} weight="bold" />
      </CloseCartButton>

      <ItemsListContainer>
        <h2>Sacola de compras</h2>
        {cartCount === 0 ? (
          <p>Seu carrinho está vazio.</p>
        ) : (
          cartItems.map((item: any) => (
            <CartItem
              key={item.id}
              product={{
                id: item.id,
                name: item.name,
                image: item.image!,
                price: item.price,
                quantity: item.quantity,
                currency: item.currency as 'BRL',
                formattedPrice: item.formattedValue,
              }}
            />
          ))
        )}
      </ItemsListContainer>

      <Footer>
        <TotalContainer>
          <div>
            <p>Quantidade</p>{' '}
            <span>
              {cartCount} {cartCount > 1 ? 'itens' : 'item'}
            </span>
          </div>
          <div>
            <p>Valor Total</p>{' '}
            <span>
              {formattedTotalPrice.replace('R$', 'R$ ').replace('.', ',')}
            </span>
          </div>
        </TotalContainer>
        <PurchaseButton
          disabled={cartCount === 0 || isCreatingCheckoutSession}
          onClick={handleBuyProduct}
        >
          Finalizar compra
        </PurchaseButton>
      </Footer>
    </CartContainer>
  )
}
