describe('Koszyk', () => {

    beforeEach(() => {
        cy.visit('http://localhost:5173')
        cy.get('.produkt-card', { timeout: 10000 }).should('have.length', 5)
    })

    it('koszyk jest pusty na starcie', () => {
        cy.get('.nav-links').contains('Koszyk').click()
        cy.contains('Twój koszyk jest pusty').should('be.visible')
    })

    
    it('dodaje produkt do koszyka', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-cart').should('contain', '1')
    })


    it('dodaje kilka produktów do koszyka', () => {
        cy.get('.produkt-btn').eq(0).click()
        cy.get('.produkt-btn').eq(1).click()
        cy.get('.nav-cart').should('contain', '2')
    })


    it('wyświetla produkt w koszyku po dodaniu', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-links').contains('Koszyk').click()
        cy.get('.koszyk-item').should('have.length', 1)
    })

    
    it('usuwa produkt z koszyka', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-links').contains('Koszyk').click()
        cy.get('.koszyk-item-usun').click()
        cy.contains('Twój koszyk jest pusty').should('be.visible')
    })


    it('wyświetla sumę w koszyku', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-links').contains('Koszyk').click()
        cy.get('.koszyk-suma').should('contain', 'zł')
    })

    it('dodanie tego samego produktu dwa razy zwiększa ilość', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-cart').should('contain', '1')
    })


    it('suma w nawigacji zgadza się z ilością produktów', () => {
        cy.get('.produkt-btn').eq(0).click()
        cy.get('.produkt-btn').eq(1).click()
        cy.get('.produkt-btn').eq(2).click()
        cy.get('.nav-cart').should('contain', '3')
    })


    it('wyświetla nazwę produktu w koszyku', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-links').contains('Koszyk').click()
        cy.get('.koszyk-item-name').should('be.visible')
    })


    it('wyświetla ilość produktu w koszyku', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-links').contains('Koszyk').click()
        cy.get('.koszyk-item-ilosc').should('contain', '1')
    })


    it('wyświetla cenę pozycji w koszyku', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-links').contains('Koszyk').click()
        cy.get('.koszyk-item-cena').should('contain', 'zł')
    })


    it('wyświetla przycisk usuń przy każdej pozycji', () => {
        cy.get('.produkt-btn').eq(0).click()
        cy.get('.produkt-btn').eq(1).click()
        cy.get('.nav-links').contains('Koszyk').click()
        cy.get('.koszyk-item-usun').should('have.length', 2)
    })


    it('po usunięciu jednego zostaje drugi produkt', () => {
        cy.get('.produkt-btn').eq(0).click()
        cy.get('.produkt-btn').eq(1).click()
        cy.get('.nav-links').contains('Koszyk').click()
        cy.get('.koszyk-item-usun').first().click()
        cy.get('.koszyk-item').should('have.length', 1)
    })


    it('przycisk przejdź do płatności jest widoczny', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-links').contains('Koszyk').click()
        cy.get('.koszyk-cta').should('be.visible')
    })


    it('przycisk przejdź do płatności przenosi na stronę płatności', () => {
        cy.get('.produkt-btn').first().click()
        cy.get('.nav-links').contains('Koszyk').click()
        cy.get('.koszyk-cta').click()
        cy.url().should('include', '/platnosci')
    })


    it('koszyk pokazuje link wróć do produktów gdy pusty', () => {
        cy.get('.nav-links').contains('Koszyk').click()
        cy.contains('Wróć do produktów').should('be.visible')
    })

})