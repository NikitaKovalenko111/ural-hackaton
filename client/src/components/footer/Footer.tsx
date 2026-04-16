import { Link } from "react-router-dom"

const Footer = () => {
    return (
        <footer className="footer">
            <div className="container">
                <p className="footer__content">© Студент и Т</p>
                <nav className="header__nav footer__nav">
                    <Link to="/">Главная</Link>
                    <Link to="/mentors">Наставники</Link>
                    <Link to="/profile">Политика</Link>
                    <Link to="https://t.me" target="_blank" rel="noreferrer">✈ Telegram</Link>
                </nav>
            </div>
        </footer>
    )
}

export default Footer