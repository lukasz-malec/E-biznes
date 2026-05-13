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

})