package myblogsearch;

import com.google.inject.Inject;
import com.huaban.analysis.jieba.WordDictionary;
import com.yahoo.language.Linguistics;
import com.yahoo.language.process.Tokenizer;
import com.yahoo.language.simple.SimpleLinguistics;
import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.nio.charset.Charset;
import java.nio.file.Paths;
import java.util.Collections;
import java.util.HashSet;
import java.util.Set;
import myblogsearch.config.SimpleChineseLinguisticsConfig;
import myblogsearch.SimpleChineseTokenizer;

public class SimpleChineseLinguistics extends SimpleLinguistics {

    // Threadsafe instances
    private final String defaultStopwordsPath = "/stopwords.txt";
    private final Set<String> stopwords;

    @Inject
    public SimpleChineseLinguistics(SimpleChineseLinguisticsConfig config) {
        stopwords = getStopwords(config.stopwordsPath());
        initWordDictionary(config.dictionaryPath());
    }

    @Override
    public Tokenizer getTokenizer() {
        return new SimpleChineseTokenizer(stopwords);
    }

    @Override
    public boolean equals(Linguistics other) { return (other instanceof SimpleChineseLinguistics); }

    private Set<String> getDefaultStopwords() {
        try {
            InputStream is = getClass().getResourceAsStream(defaultStopwordsPath);
            BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(is, Charset.forName("UTF-8")));
            Set<String> stopwords = new HashSet<>();
            String temp;
            while ((temp = bufferedReader.readLine()) != null)
                stopwords.add(temp.trim());
            bufferedReader.close();
            return Collections.unmodifiableSet(stopwords);
        }
        catch (IOException e) {
            throw new IllegalArgumentException("Failed to read stopwords file '" + defaultStopwordsPath + "'", e);
        }
    }

    private Set<String> getStopwords(String stopwordsPath) {
        if (stopwordsPath.isEmpty()) return getDefaultStopwords();
        File stopwordsFile = new File(stopwordsPath);
        try (BufferedReader bufferedReader = new BufferedReader(new FileReader(stopwordsFile))) {
            Set<String> stopwords = new HashSet<>();
            String temp;
            while ((temp = bufferedReader.readLine()) != null)
                stopwords.add(temp.trim());
            return Collections.unmodifiableSet(stopwords);
        }
        catch (IOException e) {
            throw new IllegalArgumentException("Failed to read stopwords file '" + stopwordsPath + "'", e);
        }
    }

    private void initWordDictionary(String dictionaryPath) {
        if (!dictionaryPath.isEmpty()) {
            WordDictionary.getInstance().init(new String[]{dictionaryPath});
        }
    }
}

