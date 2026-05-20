describe('Interakcje', () => {

    beforeEach(() => {
        cy.visit('http://localhost:5173')
        cy.get('.produkt-card', { timeout: 10000 }).should('have.length', 5)
    })


    
    it('przycisk dodaj jest aktywny po kliknięciu', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.produkt-btn').first().should('not.be.disabled')
    })

    it('można dodać wszystkie 5 produktów do koszyka', () => {
        cy.get('.produkt-btn').each($btn => {
            cy.wrap($btn).click()
        })
        cy.get('.nav-cart').should('contain', '5')
    })


    it('licznik koszyka w nawigacji aktualizuje się po każdym dodaniu', () => {
        cy.get('.produkt-btn').eq(0).click()
        cy.get('.nav-cart').should('contain', '1')
        cy.get('.produkt-btn').eq(1).click()
        cy.get('.nav-cart').should('contain', '2')
        cy.get('.produkt-btn').eq(2).click()
        cy.get('.nav-cart').should('contain', '3')
    })

    it('kliknięcie logo przenosi na stronę główną', () => {
        cy.get('.nav-links').contains('Koszyk').click()
        cy.get('.nav-logo').click()
        cy.url().should('eq', 'http://localhost:5173/')
    })

    it('po złożeniu zamówienia formularz nadal jest widoczny', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-links').contains('Płatności').click()
        cy.get('[name="name"]').type('Jan Kowalski')
        cy.get('[name="email"]').type('jan@test.pl')
        cy.get('[name="address"]').type('ul. Testowa 1')
        cy.get('.platnosci-btn').click()
        cy.get('.platnosci-input', { timeout: 10000 }).should('be.visible')
    })

    it('po złożeniu zamówienia można wypełnić formularz ponownie', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-links').contains('Płatności').click()
        cy.get('[name="name"]').type('Jan Kowalski')
        cy.get('[name="email"]').type('jan@test.pl')
        cy.get('[name="address"]').type('ul. Testowa 1')
        cy.get('.platnosci-btn').click()
        cy.get('[name="name"]', { timeout: 10000 }).clear().type('Anna Nowak')
        cy.get('[name="name"]').should('have.value', 'Anna Nowak')
    })

    it('nawigacja podświetla aktywną stronę', () => {
        cy.get('.nav-links').contains('Koszyk').click()
        cy.get('.nav-links a.active').should('contain', 'Koszyk')
    })


    it('przycisk wróć do produktów w pustym koszyku działa', () => {
        cy.get('.nav-links').contains('Koszyk').click()
        cy.contains('Wróć do produktów').click()
        cy.url().should('eq', 'http://localhost:5173/')
    })


    it('suma w nawigacji aktualizuje się po usunięciu produktu', () => {
        cy.get('.produkt-btn').eq(0).click()
        cy.get('.produkt-btn').eq(1).click()
        cy.get('.nav-cart').should('contain', '2')
        cy.get('.nav-links').contains('Koszyk').click()
        cy.get('.koszyk-item-usun').first().click()
        cy.get('.nav-cart').should('contain', '1')
    })


    it('przycisk złóż zamówienie jest zablokowany podczas wysyłania', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-links').contains('Płatności').click()
        cy.get('[name="name"]').type('Jan Kowalski')
        cy.get('[name="email"]').type('jan@test.pl')
        cy.get('[name="address"]').type('ul. Testowa 1')
        cy.get('.platnosci-btn').click()
        cy.get('.platnosci-btn').should('be.disabled')
    })

})