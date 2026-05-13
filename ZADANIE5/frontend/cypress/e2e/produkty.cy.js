// testy frontend dla komponentu Produkty.jsx


describe('Produkty', () => {

    beforeEach(() => {
        cy.visit('http://localhost:5173')
        cy.get('.produkt-card', { timeout: 10000 }).should('have.length', 5)
    })

    
    it('wyświetla stronę główną', () => {
        cy.get('.nav-logo').should('contain', 'SKlep React JS + GO')
    })


    it('wyświetla listę 5 produktów', () => {
        cy.get('.produkt-card').should('have.length', 5)
    })


    it('wyświetla nazwy produktów', () => {
        cy.get('.produkt-name').should('have.length', 5)
    })


    it('wyświetla ceny z symbolem zł', () => {
        cy.get('.produkt-price').each($el => {
        cy.wrap($el).should('contain', 'zł')
        })
    })


    it('wyświetla przyciski dodaj do koszyka', () => {
        cy.get('.produkt-btn').should('have.length', 5)
    })


    it('wyświetla hero section', () => {
        cy.contains('Elektronika').should('be.visible')
    })

})