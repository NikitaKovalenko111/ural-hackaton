import type { JSX } from "react";
import type React from "react";
import { Link } from "react-router-dom";

const Header: React.FC = (): JSX.Element => {
    return (
        <header className="header">
            <div className="container header__content">
                <div className="header__logo">✿ Студент и Т</div>
                <nav className="header__nav">
                    <Link to="/">Главная</Link>
                    <Link to="/events">События</Link>
                    <Link to="/mentors">Наставники</Link>
                    <Link to="/profile" className="btn btn--primary btn--sm">Профиль</Link>
                </nav>
            </div>
        </header>
    )
}

export default Header