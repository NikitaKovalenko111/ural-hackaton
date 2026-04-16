import type { JSX } from "react";
import type React from "react";
import { Link } from "react-router-dom";
import type { IUser } from "../../types";
import { useSelector } from "react-redux";
import type { RootState } from "../../redux/store";

const Header: React.FC = (): JSX.Element => {
    const user: IUser | null = useSelector((state: RootState) => state.users.user)

    return (
        <header className="header">
            <div className="container header__content">
                <div className="header__logo">✿ Студент и Т</div>
                <nav className="header__nav">
                    <Link to="/hubs">Хабы</Link>
                    <Link to="/events">События</Link>
                    <Link to="/mentors">Наставники</Link>
                    {
                        user ? (
                            <Link to="/profile" className="btn btn--primary btn--sm">
                                Профиль
                            </Link>
                        ) : (
                            <Link to="/login" className="btn btn--primary btn--sm">
                                Войти
                            </Link>
                        )
                    }
                </nav>
            </div>
        </header>
    )
}

export default Header