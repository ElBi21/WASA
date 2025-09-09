-- Dummy values
INSERT INTO Users (name, display_name, photo, bio) VALUES
('leo', 'leo', '', '''t zal wel zijn'),
('sara', 'sara', '', 'Hi, I''m a WASAText user'),
('marco', 'marco', '', 'Sup :D');

INSERT INTO Chats (isPrivate, name, description, photo) VALUES
(FALSE, 'I magici 4', 'Pazzi furiosi deh', ''),
(FALSE, 'Diabolici', 'Vamonos', ''),
(FALSE, 'Le bimbe di andrea', 'Oscar puzza', '');

INSERT INTO ChatsUsers (chat, user) VALUES
(1, 'leo'),
(2, 'leo'),
(3, 'sara'),
(2, 'sara'),
(3, 'marco');

INSERT INTO Reactions (id, message, content, user) VALUES
(1, 4, '😍', 'sara'),
(2, 4, '🗿', 'leo'),
(3, 5, '⚒️', 'leo');

-- Test for inner join
SELECT r.id, r.isPrivate, r.name, r.description, r.photo
FROM ChatsUsers l INNER JOIN Chats r ON l.chat = r.id WHERE user = 'leo';

SELECT r.name, r.display_name, r.photo, r.bio
FROM ChatsUsers l INNER JOIN Users r ON l.user = r.name WHERE l.chat = 2;