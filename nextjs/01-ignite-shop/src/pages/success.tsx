import { ImageContainer, SuccessContainer } from '@/styles/pages/success'
import Link from 'next/link'

export default function Success() {
  return (
    <SuccessContainer>
      <h1>Compra efetuada</h1>
      <ImageContainer />
      <p>
        Uhuul <strong>Miguel Castro</strong>, sua{' '}
        <strong>camiseta Berzerk</strong> já está a caminho de sua casa.
      </p>

      <Link href={'/'}>Voltar ao catalogo</Link>
    </SuccessContainer>
  )
}
