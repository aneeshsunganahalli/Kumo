CREATE TABLE users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  username TEXT UNIQUE NOT NULL,
  email TEXT UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE backgrounds (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  original_name TEXT NOT NULL,
  stored_filename TEXT UNIQUE NOT NULL,
  file_type TEXT CHECK (file_type IN ('video/mp4', 'image/png', 'image/jpeg', 'image/gif')),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE audio_tracks(
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  original_name TEXT NOT NULL,
  stored_filename TEXT UNIQUE NOT NULL,
  file_type TEXT CHECK (file_type IN ('audio/mp3', 'audio/wav')),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE ui_preferences (
  user_id INTEGER PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
  font_family TEXT DEFAULT 'Inter',
  theme_mode TEXT DEFAULT 'dark',
  fallback_color TEXT DEFAULT '#141414',
  active_background_id INTEGER REFERENCES backgrounds(id) ON DELETE SET NULL,
  active_audio_id INTEGER REFERENCES audio_tracks(id) ON DELETE SET NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE quotes (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  content TEXT NOT NULL,
  author TEXT DEFAULT 'Unknown',
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sessions (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  session_date DATE NOT NULL DEFAULT CURRENT_DATE,
  total_study_seconds INTEGER DEFAULT 0,
  notes TEXT,
  UNIQUE (user_id, session_date)
);